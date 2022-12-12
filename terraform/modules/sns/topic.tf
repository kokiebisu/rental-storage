resource "aws_sns_topic" "authentication" {
    name = "${var.environment}-${var.namespace}-authentication-topic"
}

resource "aws_sns_topic" "booking" {
    name = "${var.environment}-${var.namespace}-booking-topic"
}

resource "aws_sns_topic" "listing" {
    name = "${var.environment}-${var.namespace}-listing-topic"
}

resource "aws_sns_topic" "payment" {
    name = "${var.environment}-${var.namespace}-payment-topic"
}

resource "aws_sns_topic" "user" {
    name = "${var.environment}-${var.namespace}-user-topic"
}