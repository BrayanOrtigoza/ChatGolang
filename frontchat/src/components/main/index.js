import 'bootstrap/dist/css/bootstrap.min.css';
import React, { Component } from 'react';
import './../../styles/main.css'
import ContainerChannels from './../channels/index'

class Main extends Component {

    render() {
        return (
            <div className="container">
                <h3 className=" text-center">Chat</h3>
                <div className="messaging">
                    <ContainerChannels/>
                </div>
            </div>
            )
    }
}

export default Main;
