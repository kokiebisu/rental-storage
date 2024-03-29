output "authentication_topic_arn" {
  value = aws_sns_topic.authentication.arn
}

output "booking_topic_arn" {
  value = aws_sns_topic.booking.arn
}

output "space_topic_arn" {
  value = aws_sns_topic.space.arn
}

output "payment_topic_arn" {
  value = aws_sns_topic.payment.arn
}

output "user_topic_arn" {
  value = aws_sns_topic.user.arn
}