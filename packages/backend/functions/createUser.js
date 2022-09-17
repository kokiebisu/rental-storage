const uuidv4 = require("uuid/v4");
const Client = require("serverless-mysql");

exports.handler = async (_, obj) => {
  var client = Client({
    config: {
      host: process.env.MYSQL_HOST,
      database: process.env.DB_NAME,
      user: process.env.USERNAME,
      password: process.env.PASSWORD,
    },
  });
  await init(client);
  var userUUID = uuidv4();
  let user = await client.query("INSERT INTO users (uuid, name) VALUES(?,?)", [
    userUUID,
    obj.input.Name,
  ]);
  for (let index = 0; index < obj.input.Posts.length; index++) {
    const element = obj.input.Posts[index];
    await client.query(
      "INSERT INTO posts (uuid, text, user_id) VALUES(?, ?, ?)",
      [uuidv4(), element.Text, user.insertId]
    );
  }
  var resp = await fetchUser(client, userUUID);
  client.quit();
  return resp;
};

const init = async (client) => {
  await client.query(`
        CREATE TABLE IF NOT EXISTS users
        (
            id MEDIUMINT UNSIGNED not null AUTO_INCREMENT, 
            created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            uuid char(36) not null, 
            name varchar(100) not null, 
            PRIMARY KEY (id)
        );  
        `);
  await client.query(`
        CREATE TABLE IF NOT EXISTS posts
        (
            id MEDIUMINT UNSIGNED not null AUTO_INCREMENT, 
            created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            uuid char(36) not null, 
            text varchar(100) not null, 
            user_id MEDIUMINT UNSIGNED not null,
            PRIMARY KEY (id)
        );  
        `);
};

const fetchUser = async (client, uuid) => {
  var user = {};
  var userFromDb = await client.query(
    `
          select id, uuid, name from users where uuid = ? 
          `,
    [uuid]
  );
  if (userFromDb.length == 0) {
    return null;
  }
  var postsFromDb = await client.query(
    `
          select uuid, text from posts where user_id = ?
          `,
    [userFromDb[0].id]
  );

  user.UUID = userFromDb[0].uuid;
  user.Name = userFromDb[0].name;

  if (postsFromDb.length > 0) {
    user.Posts = postsFromDb.map(function (x) {
      return { UUID: x.uuid, Text: x.text };
    });
  }
  return user;
};
