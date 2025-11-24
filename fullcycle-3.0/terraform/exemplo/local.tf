resource "local_file" "exemplo" {
  filename = "teste.txt"
  content  = var.conteudo
}

variable "conteudo" {
  type = string
  description = "Conteúdo a ser escrito no arquivo local"
  default = "value padrão se não for fornecido"
}

output "id_arquivo" {
  value = local_file.exemplo.id
}

output "variavel-conteudo" {
  value = var.conteudo
}

output "conteudo_arquivo" {
  value = data.local_file.exemplo-data.content
}

data "local_file" "exemplo-data" {
  filename = local_file.exemplo.filename
}