package geodataupdate

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadAndVerify(url, checksumURL, destPath string) error {
	// Download data file
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download data file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status code: %d", resp.StatusCode)
	}

	// Create temporary file
	tmpFile, err := ioutil.TempFile("", "geodata_*.dat")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer tmpFile.Close()

	// Download data to temporary file
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to download data to temporary file: %w", err)
	}

	// Verify checksum
	err = verifyChecksum(tmpFile.Name(), checksumURL)
	if err != nil {
		return fmt.Errorf("checksum verification failed: %w", err)
	}

	// Move temporary file to destination
	err = os.Rename(tmpFile.Name(), destPath)
	if err != nil {
		return fmt.Errorf("failed to move temporary file to destination: %w", err)
	}

	return nil
}

func verifyChecksum(filePath, checksumURL string) error {
	// Download checksum file
	resp, err := http.Get(checksumURL)
	if err != nil {
		return fmt.Errorf("failed to download checksum file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("checksum download failed with status code: %d", resp.StatusCode)
	}

	// Read checksum
	checksumData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read checksum data: %w", err)
	}
	expectedChecksum := string(checksumData)

	// Calculate actual checksum
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open data file for checksum: %w", err)
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return fmt.Errorf("failed to calculate checksum: %w", err)
	}
	actualChecksum := fmt.Sprintf("%x", hash.Sum(nil))

	// Compare checksums
	if actualChecksum != expectedChecksum {
		return fmt.Errorf("checksum mismatch: expected %s, got %s", expectedChecksum, actualChecksum)
	}

	return nil
}
