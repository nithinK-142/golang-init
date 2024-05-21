import express, { json, urlencoded } from "express";

const app = express();
const PORT = 8000;

app.use(json());
app.use(urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("server");
});

app.get("/get", (req, res) => {
  res.status(200).json({ message: "Server saying hello" });
});

app.post("/post", (req, res) => {
  let jsonObj = req.body; // your JSON
  res.status(200).send(jsonObj);
});

app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => {
  console.log(`server started http://localhost:${PORT}`);
});
