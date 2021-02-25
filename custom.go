package logrotate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	customBackupTimeFormat = "20060102150405.000"
	customCompressPath     = "/zip"
)

// customBackupName creates a new filename from the given name, inserting a timestamp
// between the filename and the extension, using the local time if requested
// (otherwise UTC).
func customBackupName(name string, local bool) (string, error) {
	dir := customFileDir(filepath.Dir(name))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("can't make directories for backup logfile: %s", err)
	}

	filename := filepath.Base(name)

	t := currentTime()
	if !local {
		t = t.UTC()
	}

	timestamp := strings.Replace(t.Format(customBackupTimeFormat), ".", "", -1)
	filename = regexp.MustCompile(`-[\w]+-\[`).ReplaceAllString(filename, "-"+timestamp+"-[")
	return filepath.Join(dir, filename), nil
}

func customNewName(name string, local bool) string {
	dir := filepath.Dir(name)
	filename := filepath.Base(name)

	t := currentTime()
	if !local {
		t = t.UTC()
	}
	timestamp := strings.Replace(t.Format(customBackupTimeFormat), ".", "", -1)
	filename = regexp.MustCompile(`-[\w]+-\[`).ReplaceAllString(filename, "-"+timestamp+"-[")
	return filepath.Join(dir, filename)
}

// customTimeFromName extracts the formatted time from the filename by stripping off
// the filename's prefix and extension. This prevents someone's filename from
// confusing time.parse.
func customTimeFromName(filename string) (time.Time, error) {
	if !strings.HasSuffix(filename, ".log") && !strings.HasSuffix(filename, ".log.gz") {
		return time.Time{}, errors.New("mismatched extension")
	}
	timetmp := strings.Trim(regexp.MustCompile(`-[\w]+-\[`).FindString(filename), "-[")
	if timetmp == "" {
		return time.Time{}, errors.New("mismatched extension")
	}
	timelen := len(timetmp)
	timetmp = timetmp[:timelen-3] + "." + timetmp[timelen-3:]
	return time.Parse(customBackupTimeFormat, timetmp)
}

func customFileDir(dir string) string {
	return dir + customCompressPath
}
