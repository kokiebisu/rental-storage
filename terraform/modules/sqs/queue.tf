resource "aws_sqs_queue" "booking" {
  name = "${var.environment}-${var.namespace}-booking-queue"
  depends_on = [
    aws_sqs_queue.booking_dead_letter
  ]
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.booking_dead_letter.arn
    maxReceiveCount     = 3
  })
}

resource "aws_sqs_queue" "booking_dead_letter" {
  name = "${var.environment}-${var.namespace}-booking-dead-letter-queue"
}

resource "aws_sqs_queue" "space" {
  name = "${var.environment}-${var.namespace}-space-queue"
  depends_on = [
    aws_sqs_queue.space_dead_letter
  ]
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.space_dead_letter.arn
    maxReceiveCount     = 3
  })
}

resource "aws_sqs_queue" "space_dead_letter" {
  name = "${var.environment}-${var.namespace}-space-dead-letter-queue"
}

resource "aws_sqs_queue" "payment" {
  name = "${var.environment}-${var.namespace}-payment-queue"
  depends_on = [
    aws_sqs_queue.payment_dead_letter
  ]
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.payment_dead_letter.arn
    maxReceiveCount     = 3
  })
}

resource "aws_sqs_queue" "payment_dead_letter" {
  name = "${var.environment}-${var.namespace}-payment-dead-letter-queue"
}

resource "aws_sqs_queue" "user" {
  name = "${var.environment}-${var.namespace}-user-queue"
  depends_on = [
    aws_sqs_queue.user_dead_letter
  ]
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.user_dead_letter.arn
    maxReceiveCount     = 3
  })
}

resource "aws_sqs_queue" "user_dead_letter" {
  name = "${var.environment}-${var.namespace}-user-dead-letter-queue"
}

