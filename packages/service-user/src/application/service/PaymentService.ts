import axios from "axios";
import { PaymentRepositoryImpl } from "../../Adapter/Repository";
import { Payment } from "../../Domain/Model";
import { LoggerUtil } from "../../Utils";
import { AddPaymentInput, PaymentRepository } from "../Port";
import { PaymentService } from "../Port";

export class PaymentServiceImpl implements PaymentService {
  private _paymentRepository: PaymentRepository;

  private _logger: LoggerUtil;

  private constructor(paymentRepository: PaymentRepository) {
    this._paymentRepository = paymentRepository;
    this._logger = new LoggerUtil("PaymentServiceImpl");
  }

  public static async create() {
    const paymentRepository = await PaymentRepositoryImpl.create();
    await paymentRepository.setup();
    return new PaymentServiceImpl(paymentRepository);
  }

  public async addPayment(data: AddPaymentInput) {
    this._logger.info(data, "addPayment()");
    // get stripe customer id for the user
    try {
      const response: { data: { providerId: string; providerType: string } } =
        await axios.post(
          `${process.env.SERVICE_API_ENDPOINT}/payments/customer`,
          {
            emailAddress: data.emailAddress,
            firstName: data.firstName,
            lastName: data.lastName,
          }
        );

      const payment = new Payment({
        userId: data.userId,
        providerType: response.data.providerType,
        providerId: response.data.providerId,
      });
      await this._paymentRepository.save(payment);
    } catch (err) {
      console.error(err);
    }
  }
}
