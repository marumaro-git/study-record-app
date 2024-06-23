terraform {
  backend "s3" {
    bucket = "study-record-terraform"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
  }
}
