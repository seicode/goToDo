import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:1323";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: []
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { task } = this.state;
    if (task) {
      axios
        .post(
          endpoint + "/api/task",
          {
            task
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        )
        .then(res => {
          this.getTask();
          this.setState({
            task: ""
          });
          console.log("onSubmit:", res.statusText);
        });
    }
  };

  getTask = () => {
    axios.get(endpoint + "/api/task").then(res => {
      let str = ""
      if (res.data) {
        // remove unexpected string from data
        if (res.data.indexOf(`\nnull\n`) !== -1) {
          str = res.data.replace(`\nnull\n`, "");
        } else {
          str = res.data
        }
        let obj = JSON.parse(str);
        console.log("No. of items:", obj.length);
        this.setState({
          items: obj.map((item) => {
            // status color => Done:green ToDo:yellow
            let color = "yellow";
            if (item.status) {
              color = "green";
            }
            return (
              <Card key={item._id} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word" }}>{item.task}</div>
                  </Card.Header>

                  <Card.Meta textAlign="right">
                    <span onClick={() => this.updateTask(item._id)}>
                      <Icon
                        name="check circle"
                        color="green"
                      />
                      <span style={{ paddingRight: 10 }}>Done</span>
                    </span>
                    <span onClick={() => this.undoTask(item._id)}>
                      <Icon
                        name="undo"
                        color="yellow"
                      />
                      <span style={{ paddingRight: 10 }}>Undo</span>
                    </span>
                    <span onClick={() => this.deleteTask(item._id)}>
                      <Icon
                        name="delete"
                        color="red"
                      />
                      <span style={{ paddingRight: 10 }}>Delete</span>
                    </span>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          })
        });
      } else {
        this.setState({
          items: []
        });
      }
    });
  };

  updateTask = id => {
    axios
      .put(endpoint + "/api/task/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log("updateTask", res.statusText);
        this.getTask();
      });
  };

  undoTask = id => {
    axios
      .put(endpoint + "/api/undoTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log("undoTask", res.status);
        this.getTask();
      });
  };

  deleteTask = id => {
    axios
      .delete(endpoint + "/api/deleteTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log("deleteTask", res.statusText);
        this.getTask();
      });
  };
  render() {
    return (
      <div>
        <div className="row" style={{marginTop: "20px", marginBottom: "20px"}}>
          <Header className="header" as="h2">
            GoToDo List
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create Task"
              style={{marginBottom: "20px"}}
            />
            {/* <Button >Create Task</Button> */}
          </Form>
        </div>
        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDoList;