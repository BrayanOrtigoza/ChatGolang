import 'bootstrap/dist/css/bootstrap.min.css';
import React, { Component } from 'react';
import './../../styles/main.css'
import ContainerChannels from './../channels/index'

class Main extends Component {


    render() {
        return (
            <div className="container">
                <input type="submit" className="fadeIn fourth" value="Salir" onClick={()=> this.props.logout()}/>
                <h3 className=" text-center">Chat</h3>
                <div className="messaging">
                    <ContainerChannels logout={this.props.logout.bind(this)}/>
                </div>
            </div>
            )
    }
}

export default Main;
