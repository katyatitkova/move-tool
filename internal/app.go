package internal

import (
	"fmt"
	"github.com/alexfedosov/move-tool/internal/ablmodels"
	"github.com/brianvoe/gofakeit/v7"
	"strings"
)

func sanitizePresetName(presetName string) string {
	var result strings.Builder

	for _, char := range presetName {
		if (char >= 'a' && char <= 'z') || char == '_' {
			result.WriteRune(char)
		} else {
			result.WriteRune('_')
		}
	}

	return result.String()
}

func SliceSampleIntoDrumRack(inputFilePath string, outputFolderPath string, numberOfSlices int, customPresetName string) (err error) {
	err = gofakeit.Seed(0)
	if err != nil {
		return err
	}

	var presetName string
	if customPresetName != "" {
		presetName = customPresetName
	} else {
		presetName = strings.ToLower(fmt.Sprintf("%s_%s", gofakeit.HipsterWord(), gofakeit.AdverbPlace()))
		presetName = sanitizePresetName(presetName)
	}

	presetFolderPath, err := createFolderIfNotExist(outputFolderPath, presetName)
	if err != nil {
		return err
	}
	samplesFolderPath, err := createFolderIfNotExist(presetFolderPath, "Samples")
	if err != nil {
		return err
	}
	samples, err := writeAudioFileSlices(inputFilePath, samplesFolderPath, numberOfSlices, presetName)
	if err != nil {
		return err
	}

	preset := ablmodels.NewDrumRackDevicePresetWithSamples(*samples)

	err = writePresetFile(preset, presetFolderPath)
	if err != nil {
		return err
	}
	err = archivePresetBundle(presetName, presetFolderPath, outputFolderPath)
	if err != nil {
		return err
	}

	err = removeDirectory(presetFolderPath)
	if err != nil {
		return err
	}

	return nil
}
