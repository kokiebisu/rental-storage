resource "aws_sns_topic_subscription" "booking" {
  depends_on = [
    aws_sqs_queue.payment,
    var.user_topic_arn
  ]
  protocol = "sqs"
  topic_arn = "${var.user_topic_arn}"
  endpoint = "${aws_sqs_queue.booking.arn}"
}