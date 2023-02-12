output "booking_queue_arn" {
  value = aws_sqs_queue.booking.arn
}

output "space_queue_arn" {
    value = aws_sqs_queue.space.arn
}

output "payment_queue_arn" {
    value = aws_sqs_queue.payment.arn
}

output "user_queue_arn" {
    value = aws_sqs_queue.user.arn
}