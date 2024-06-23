variable "eks_allow_ports" {
  description = "List of ports to allow"
  type        = list(number)
  default     = [9443]
}
