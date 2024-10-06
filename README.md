# Botredirect

O script verifica se existe um arquivo com o prefixo _fulldomains200.txt, caso contrário, pede para você informar o nome do arquivo. Em seguida, para cada linha do arquivo, executa o comando descrito e exibe os resultados na tela.

# Instalação

```
go install https://github.com/lupedsagaces/botredirect@latest
```

**após vários testes, o oneliner ficou assim:**
```bash
echo exemplo.com | alterx -enrich | waybackurls | dnsx | httpx -silent -mc 302 | grep -a -i \=http | qsreplace 'http://evil.com' | while read host do;do curl -s -L $host -I|grep “evil.com” && echo -e “$host \033[0;31mVulnerable\n” ;done
```


1. **`echo exemplo.com`**: Aqui, o domínio `exemplo.com` será uma lista já validada de subdomínios.
    
2. **`alterx -enrich`**: O comando `alterx` está sendo usado com a opção `-enrich`, que enriquece a wordlist gerando novas variações de subdomínios baseadas em padrões e dados do domínio fornecido. Ele pode criar permutações novas e mais focadas em cenários específicos.
    
3. **`gau`**: A ferramenta coleta URLs arquivadas associadas ao domínio fornecido, extraídas da Wayback Machine, entre outras fontes. Isso ajuda a obter uma lista de URLs que historicamente existiram para o domínio.
    
4. **`dnsx`**: Verifica se as URLs geradas ainda possuem registros DNS válidos. Ele elimina URLs que não possuem resolução DNS.
    
5. **`httpx -silent -mc 302`**: Envia requisições HTTP para verificar o código de resposta. O `-mc 302` filtra URLs que retornam um código de redirecionamento (302).
    
6. **`grep -a -i \=http`**: Procura por parâmetros em URLs que contenham `=http`. Está buscando links embutidos nos parâmetros, comuns em redirecionamentos.
    
7. **`qsreplace 'http://evil.com'`**: Substitui qualquer valor do parâmetro encontrado por `http://evil.com`, simulando a tentativa de inserção de uma URL maliciosa para verificar possíveis vulnerabilidades de redirecionamento aberto (open redirect).
    
8. **`while read host do; do curl -s -L $host -I`**: Para cada URL gerada, `curl` faz uma requisição HTTP para seguir redirecionamentos (`-L`) e captura apenas os cabeçalhos HTTP (`-I`).
    
9. **`grep “evil.com”`**: Verifica se o cabeçalho HTTP retornado contém a URL maliciosa `evil.com`, sugerindo que a URL pode estar vulnerável a redirecionamento.
    
10. **`echo -e “$host \033[0;31mVulnerable\n”`**: Se `evil.com` for encontrado, ele imprime a URL vulnerável com a mensagem "Vulnerable" em vermelho.
    
