# Cria uma VPC (Virtual Private Cloud) com o bloco CIDR especificado
resource "aws_vpc" "vpc-fullcycle" {
  cidr_block = "10.0.0.0/16" # Define o intervalo de IPs da VPC
  tags = {
    Name = "${var.prefix}-vpc" # Nome da VPC para identificação
  }
}

# Obtém as zonas de disponibilidade disponíveis na região configurada
data "aws_availability_zones" "available" {}

# Cria duas subnets públicas dentro da VPC, uma em cada zona de disponibilidade
resource "aws_subnet" "subnets" {
  count = 2 # Cria duas subnets
  availability_zone = data.aws_availability_zones.available.names[count.index] # Define a zona de disponibilidade
  vpc_id = aws_vpc.vpc-fullcycle.id # Associa a subnet à VPC criada
  cidr_block = "10.0.${count.index}.0/24" # Define o intervalo de IPs para cada subnet
  map_public_ip_on_launch = true # Permite que instâncias na subnet recebam IPs públicos automaticamente
  tags = {
    Name = "${var.prefix}-subnet-${count.index}" # Nome da subnet para identificação
  }
}

# Cria um Internet Gateway para permitir comunicação da VPC com a internet
resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.vpc-fullcycle.id # Associa o IGW à VPC criada
  tags = {
    Name = "${var.prefix}-igw" # Nome do IGW para identificação
  }
}

# Cria uma tabela de rotas para gerenciar o tráfego da VPC
resource "aws_route_table" "route-table" {
  vpc_id = aws_vpc.vpc-fullcycle.id # Associa a tabela de rotas à VPC criada
  route {
    cidr_block = "0.0.0.0/0" # Define uma rota padrão para todo o tráfego
    gateway_id = aws_internet_gateway.igw.id # Direciona o tráfego para o Internet Gateway
  }
  tags = {
    Name = "${var.prefix}-route-table" # Nome da tabela de rotas para identificação
  }
}

# Associa a tabela de rotas criada às subnets públicas
resource "aws_route_table_association" "new-rtb-association" {
  count = 2 # Cria associações para as duas subnets
  subnet_id = aws_subnet.subnets[count.index].id # Associa a tabela de rotas à subnet correspondente
  route_table_id = aws_route_table.route-table.id # Especifica a tabela de rotas a ser associada
}