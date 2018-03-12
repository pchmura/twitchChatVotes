import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
    constructor(props){
      super(props)
      this.state = {
        message:"this shouldnt be here",
        ws: null,
        channel: "",
        option1: "",
        option2: "",
        emote1: "",
        emote2: "",
        duration: 0,
        votes1: 0,
        votes2: 0
      }

      this.sendMessage = this.sendMessage.bind(this);
      this.sendMessage2 = this.sendMessage2.bind(this);
      this.sendFormData = this.sendFormData.bind(this);
      this.handleInputChange = this.handleInputChange.bind(this);
    }

  componentDidMount(){
    let ws = new WebSocket('ws://localhost:8000/ws');
    
    ws.addEventListener('message', (e) =>{
      let message = JSON.parse(e.data);
      console.log(message);
      this.setState({message:message.data});
    })
    this.setState({ws: ws});
  }
  
  sendMessage(){
    this.state.ws.send(
      JSON.stringify({
          data: "pls work some more fam",
          type: "testMessage"
      }
  ));
  }

  sendMessage2(){
    this.state.ws.send(
      JSON.stringify({
          data: "pls work some more fam",
          type: "testMessage2"
      }
  ));
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  sendFormData(event){
    
    console.log(this.state)
    event.preventDefault();
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">{this.state.message}</h1>
        </header>
        <p className="App-intro">
          <button onClick={this.sendMessage}>Testing Case Button 1</button>
          <button onClick={this.sendMessage2}>Testing Case Button 2</button>
        </p>
        <div>
        <form onSubmit={this.sendFormData}>
          <label>
            Vote Option 1:
            <input type="text" name="option1" value={this.state.option1} onChange={this.handleInputChange} />
          </label>
          <br/>
          <label>
            Vote Option 2:
            <input type="text" name="option2" value={this.state.option2} onChange={this.handleInputChange} />
          </label>
          <br/>
          <label>
            Vote Emote 1:
            <input type="text" name="emote1" value={this.state.emote1} onChange={this.handleInputChange} />
          </label>
          <br/>
          <label>
            Vote Emote 2:
            <input type="text" name="emote2" value={this.state.emote2} onChange={this.handleInputChange} />
          </label>
          <br/>
          <label>
            Channel:
            <input type="text" name="channel" value={this.state.channel} onChange={this.handleInputChange} />
          </label>
          <br/>
          <input type="submit" value="Submit" />
        </form>
        </div>
      </div>
    );
  }
}

export default App;
