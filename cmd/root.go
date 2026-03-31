package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/natori/nasa-apod-cli/internal/apod"
	"github.com/natori/nasa-apod-cli/internal/ascii"
	"github.com/natori/nasa-apod-cli/internal/translate"
	"github.com/spf13/cobra"
)

var (
	flagDate   string
	flagRandom bool
	flagJa     bool
	flagASCII  bool
)

var rootCmd = &cobra.Command{
	Use:   "apod",
	Short: "NASA Astronomy Picture of the Day CLI",
	Long:  "Fetch and display NASA's Astronomy Picture of the Day from the terminal.",
	RunE:  run,
}

func init() {
	rootCmd.Flags().StringVar(&flagDate, "date", "", "Fetch APOD for a specific date (YYYY-MM-DD)")
	rootCmd.Flags().BoolVar(&flagRandom, "random", false, "Fetch a random APOD")
	rootCmd.Flags().BoolVar(&flagJa, "ja", false, "Translate explanation to Japanese")
	rootCmd.Flags().BoolVar(&flagASCII, "ascii", false, "Display image as ASCII art in terminal")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var datePattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

func run(cmd *cobra.Command, args []string) error {
	_ = godotenv.Load()

	apiKey := os.Getenv("NASA_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("NASA_API_KEY is not set. Set it as an environment variable or in a .env file")
	}

	if flagDate != "" && !datePattern.MatchString(flagDate) {
		return fmt.Errorf("invalid date format: %s (expected YYYY-MM-DD)", flagDate)
	}

	client := apod.NewClient(apiKey)

	var result *apod.Response
	var err error

	switch {
	case flagRandom:
		result, err = client.GetRandom()
	case flagDate != "":
		result, err = client.GetByDate(flagDate)
	default:
		result, err = client.GetToday()
	}
	if err != nil {
		return fmt.Errorf("failed to fetch APOD: %w", err)
	}

	printResult(result)

	if flagJa {
		translated, err := translate.ToJapanese(result.Explanation)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: translation failed: %v\n", err)
		} else {
			fmt.Println("\n--- 日本語訳 ---")
			fmt.Println(translated)
		}
	}

	if flagASCII {
		if result.MediaType != "image" {
			fmt.Println("\nNote: ASCII art is only available for image type APOD (this is a video).")
			return nil
		}
		art, err := ascii.RenderFromURL(result.URL)
		if err != nil {
			return fmt.Errorf("failed to render ASCII art: %w", err)
		}
		fmt.Println()
		fmt.Println(art)
	}

	return nil
}

func printResult(r *apod.Response) {
	fmt.Printf("Title: %s\n", r.Title)
	fmt.Printf("Date:  %s\n", r.Date)
	if r.Copyright != "" {
		fmt.Printf("Copyright: %s\n", r.Copyright)
	}
	fmt.Printf("URL:   %s\n", r.URL)
	if r.HDURL != "" {
		fmt.Printf("HD:    %s\n", r.HDURL)
	}
	fmt.Println()
	fmt.Println("--- Explanation ---")
	fmt.Println(r.Explanation)
}
