import React, { useEffect } from "react";
import { Request } from "./newspb/protobuf/news_pb";
import { NewServiceClient } from "./newspb/protobuf/news_grpc_web_pb";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import Avatar from "@mui/material/Avatar";
import Typography from "@mui/material/Typography";
import Card from "@mui/material/Card";
import CardHeader from "@mui/material/CardHeader";
import CardMedia from "@mui/material/CardMedia";
import CardContent from "@mui/material/CardContent";
import CardActions from "@mui/material/CardActions";
import Link from "@mui/material/Link";
import { red } from "@mui/material/colors";
import NotFound from "./notfound.gif";

var client = new NewServiceClient("http://localhost:8000", null, null);

function App() {
  const [newsList, setNewsList] = React.useState([]);
  const [type, setType] = React.useState("top");
  const [page, setPage] = React.useState(1);
  const [size, setSize] = React.useState(10);

  const handleTypeChange = (event) => {
    setType(event.target.value);
    console.log(event.target.value);
    getHotNews(event.target.value, page, size);
  };

  const handleSizeChange = (event) => {
    setSize(event.target.value);
    console.log(event.target.value);
    getHotNews(type, page, event.target.value);
  };

  const handlePageChange = (event) => {
    setPage(event.target.value);
    console.log(event.target.value);
    getHotNews(type, event.target.value, size);
  };

  const getHotNews = (type, page, size) => {
    console.log(type, page, size);
    var request = new Request();
    request.setType(type);
    request.setPage(page);
    request.setSize(size);
    client.getHotTopNews(request, {}, (error, reply) => {
      if (!error) {
        setNewsList(reply.getNewsList());
      } else {
        console.log(error);
      }
    });
  };

  useEffect(() => {
    getHotNews(type, page, size);
  }, [type, page, size]);

  return (
    <Container>
      <Box>
        <FormControl sx={{ m: 1, minWidth: 120 }}>
          <InputLabel htmlFor="type-select">Type</InputLabel>
          <Select
            defaultValue="top"
            id="type-select"
            label="Type"
            value={type}
            onChange={handleTypeChange}
          >
            <MenuItem value={"top"}>??????</MenuItem>
            <MenuItem value={"guonei"}>??????</MenuItem>
            <MenuItem value={"guoji"}>??????</MenuItem>
            <MenuItem value={"yule"}>??????</MenuItem>
            <MenuItem value={"tiyu"}>??????</MenuItem>
            <MenuItem value={"junshi"}>??????</MenuItem>
            <MenuItem value={"keji"}>??????</MenuItem>
            <MenuItem value={"caijing"}>??????</MenuItem>
            <MenuItem value={"youxi"}>??????</MenuItem>
            <MenuItem value={"qiche"}>??????</MenuItem>
            <MenuItem value={"jiankang"}>??????</MenuItem>
          </Select>
        </FormControl>
        <FormControl sx={{ m: 1, minWidth: 120 }}>
          <TextField
            id="page-select"
            label="Page"
            variant="outlined"
            value={page}
            onChange={handlePageChange}
          />
        </FormControl>
        <FormControl sx={{ m: 1, minWidth: 120 }}>
          <InputLabel htmlFor="size-select">Size</InputLabel>
          <Select
            defaultValue="5"
            id="size-select"
            label="Size"
            value={size}
            onChange={handleSizeChange}
          >
            <MenuItem value={5}>5</MenuItem>
            <MenuItem value={10}>10</MenuItem>
            <MenuItem value={20}>20</MenuItem>
            <MenuItem value={30}>30</MenuItem>
          </Select>
        </FormControl>
      </Box>
      <Box>
        <Grid
          container
          spacing={{ xs: 2, md: 3 }}
          columns={{ xs: 4, sm: 8, md: 12 }}
        >
          {newsList.map((value, index) => (
            <Grid item xs={2} sm={4} md={4}>
              <Card>
                <CardHeader
                  avatar={
                    <Avatar sx={{ bgcolor: red[500] }} aria-label="recipe">
                      {value.array[4][0]}
                    </Avatar>
                  }
                  title={value.array[4] + value.array[3]}
                  subheader={value.array[2]}
                />
                {value.array[6] === null ||
                value.array[6] === undefined ||
                value.array[6] === "" ? (
                  <CardMedia
                    component="img"
                    height="194"
                    image={NotFound}
                    alt="News cover"
                  />
                ) : (
                  <CardMedia
                    component="img"
                    height="194"
                    image={value.array[6]}
                    alt="News cover"
                  />
                )}
                <CardContent>
                  <Typography variant="body2" color="text.secondary">
                    {value.array[1]}
                  </Typography>
                </CardContent>
                <CardActions>
                  <Link
                    href={value.array[5]}
                    underline="none"
                    target="_blank"
                    rel="noopener"
                  >
                    ????????????
                  </Link>
                </CardActions>
              </Card>
            </Grid>
          ))}
        </Grid>
      </Box>
    </Container>
  );
}

export default App;
