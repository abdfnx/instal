package installer

import (
	"fmt"
	"time"
	"runtime"
	"net/http"
	"io/ioutil"

	"github.com/abdfnx/gosh"
	"github.com/briandowns/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/abdfnx/instal/core/options"
)

func RunInstal(opts *options.InstalOptions, isCLI bool, url, shell, isHidden string) error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 🔗 Installing..."
	s.Start()

	term := ""

	if runtime.GOOS == "windows" {
		term = "powershell.exe"
	} else {
		term = "bash"
	}

	if isCLI {
		if opts.Shell == "" {
			opts.Shell = term
		}
	} else {
		opts.Shell = shell
		opts.URL = url
		
		if isHidden == "true" || isHidden == "y" || isHidden == "yes" {
			opts.IsHidden = true
		} else {
			opts.IsHidden = false
		}
	}

	res, err := http.Get(opts.URL)
	
	if err != nil {
		return err
	}
	
	defer res.Body.Close()

	body, berr := ioutil.ReadAll(res.Body)

	if berr != nil {
		return berr
	}

	err, out, errout := gosh.Exec(opts.Shell, string(body))

	if err != nil {
		if !isCLI {
			fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(err.Error()).String())
			fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(errout).String())
		} else {
			fmt.Println(err)
			fmt.Println(errout)
		}
	}

	s.Stop()

	if !opts.IsHidden {
		if !isCLI {
			fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(out).String())
		} else {
			fmt.Println(out)
		}
	}

	return nil
}
