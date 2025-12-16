package service

import (
	"bufio"
	"os"
	"strings"
)

// Producer интерфейс для поставщика данных
type Producer interface {
	Produce() ([]string, error)
}

// Presenter интерфейс для вывода результатов
type Presenter interface {
	Present([]string) error
}

// Service основной сервис
type Service struct {
	prod Producer
	pres Presenter
}

// NewService создает новый сервис
func NewService(prod Producer, pres Presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}

// Run запускает обработку данных
func (s *Service) Run() error {
	// Получаем данные
	data, err := s.prod.Produce()
	if err != nil {
		return err
	}

	// Обрабатываем данные
	result := s.MaskData(data)

	// Выводим результат
	return s.pres.Present(result)
}

// MaskData маскирует данные
func (s *Service) MaskData(data []string) []string {
	var result []string

	for _, line := range data {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		words := strings.Fields(line)
		maskedWords := make([]string, len(words))

		for i, word := range words {
			if len(word) > 1 {
				maskedWords[i] = string(word[0]) + strings.Repeat("*", len(word)-1)
			} else {
				maskedWords[i] = word
			}
		}

		result = append(result, strings.Join(maskedWords, " "))
	}

	return result
}

// FileProducer читает данные из файла
type FileProducer struct {
	FilePath string
}

// NewFileProducer создает новый FileProducer
func NewFileProducer(filePath string) *FileProducer {
	return &FileProducer{FilePath: filePath}
}

// Produce читает строки из файла
func (fp *FileProducer) Produce() ([]string, error) {
	file, err := os.Open(fp.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// FilePresenter записывает данные в файл
type FilePresenter struct {
	FilePath string
}

// NewFilePresenter создает новый FilePresenter
func NewFilePresenter(filePath string) *FilePresenter {
	return &FilePresenter{FilePath: filePath}
}

// Present записывает данные в файл
func (fp *FilePresenter) Present(data []string) error {
	content := strings.Join(data, "\n")

	file, err := os.Create(fp.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
