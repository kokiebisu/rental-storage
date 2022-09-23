import { UserServiceImpl } from "../../../service";

export const handler = async (event: any) => {
  const { firstName, lastName } = event.arguments;
  const service = await UserServiceImpl.create();
  return await service.registerGuest({ firstName, lastName });

  // const { name } = event.arguments;
  // var client = Client({
  //   config: {
  //     host: process.env.MYSQL_HOST,
  //     database: process.env.DB_NAME,
  //     user: process.env.USERNAME,
  //     password: process.env.PASSWORD,
  //   },
  // });
  // await init(client);
  // var userUUID = uuidv4();
  // await client.query("INSERT INTO users (uuid, name) VALUES(?,?)", [
  //   userUUID,
  //   name,
  // ]);
  // // for (let index = 0; index < obj.input.Posts.length; index++) {
  // //   const element = obj.input.Posts[index];
  // //   await client.query(
  // //     "INSERT INTO posts (uuid, text, user_id) VALUES(?, ?, ?)",
  // //     [uuidv4(), element.Text, user.insertId]
  // //   );
  // // }
  // var resp = await fetchUser(client, userUUID);
  // client.quit();
  // return resp;
};
