// package main

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"fmt"
// 	"image"
// 	"image/png"
// 	"os"
// 	"path/filepath"

// 	"golang.org/x/sys/windows"
// )

// func getWindowsExecutableIcon(execPath string) (string, error) {
// 	// Ensure absolute path
// 	absPath, err := filepath.Abs(execPath)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Load the icons
// 	icon, err := windows.ExtractIcon(0, windows.StringToUTF16Ptr(absPath), 0)
// 	if err != nil || icon == 0 {
// 		return "", fmt.Errorf("could not extract icon for %s", absPath)
// 	}

// 	// Create icon info struct
// 	var iconInfo windows.IconInfo
// 	if err := windows.GetIconInfo(icon, &iconInfo); err != nil {
// 		return "", err
// 	}

// 	// Read bitmap data
// 	bitmapPath := fmt.Sprintf("%s.bmp", absPath)
// 	if err := windows.SaveBitmapToFile(iconInfo.HbmColor, bitmapPath); err != nil {
// 		return "", err
// 	}
// 	defer os.Remove(bitmapPath)

// 	// Read bitmap file
// 	iconData, err := os.ReadFile(bitmapPath)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Base64 encode
// 	return base64.StdEncoding.EncodeToString(iconData), nil
// }

// func convertIconToBase64(icon image.Image) (string, error) {
// 	// Encode the icon image to a PNG format
// 	var buf bytes.Buffer
// 	err := png.Encode(&buf, icon)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to encode image to PNG: %v", err)
// 	}

// 	// Encode the PNG as a base64 string
// 	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

//		// Return the base64 string
//		return base64Str, nil
//	}
package main
