package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func printBanner() {
	banner := `
	

____    ___   ______  ____     ___  ___    ____  ____     ___    __ ______ 
|    \  /   \ |      ||    \   /  _]|   \  |    ||    \   /  _]  /  ]      |
|  o  )|     ||      ||  D  ) /  [_ |    \  |  | |  D  ) /  [_  /  /|      |
|     ||  O  ||_|  |_||    / |    _]|  D  | |  | |    / |    _]/  / |_|  |_|
|  O  ||     |  |  |  |    \ |   [_ |     | |  | |    \ |   [_/   \_  |  |  
|     ||     |  |  |  |  .  \|     ||     | |  | |  .  \|     \     | |  |  
|_____| \___/   |__|  |__|\_||_____||_____||____||__|\_||_____|\____| |__|  
																				
		Open redirect search tool
				By: lupedsagaces	
	`
	fmt.Println(banner)
}

func main() {
	printBanner()
	// Define o prefixo do arquivo
	filePrefix := "_fulldomains200.txt"

	// Procura o arquivo na pasta atual
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("Erro ao ler o diretório: %v", err)
	}

	var filename string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), filePrefix) {
			filename = file.Name()
			break
		}
	}

	// Caso o arquivo não seja encontrado
	if filename == "" {
		fmt.Println("Arquivo de subdomínios não encontrado. Por favor, informe o nome do arquivo manualmente:")
		fmt.Scanln(&filename)
	}

	// Abre o arquivo para leitura
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	// Lê o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		runCommand(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}
}

// Função que executa o comando para cada domínio
func runCommand(domain string) {
	command := fmt.Sprintf(`echo %s | alterx -enrich | gau | dnsx | httpx -silent | grep -a -i =http | qsreplace 'http://evil.com' | while read host; do curl -s -L $host -I | grep 'evil.com' && echo -e "$host \033[0;31mVulnerável\n"; done`, domain)

	cmd := exec.Command("bash", "-c", command)

	// Mostra a saída do comando na tela
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Erro ao executar o comando para %s: %v", domain, err)
	}
}
