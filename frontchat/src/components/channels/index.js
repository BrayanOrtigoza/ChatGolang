import React, { Component } from 'react';
import Message from './../message/index'
import User from './../user/index'
import {getService} from "./../../services/services";
import {Routes} from "./../../services/routes";
import {postService} from "../../services/services";
import Socket from "../socket/socket";
import $ from "jquery";

class ContainerChannels extends Component {

    constructor(props) {
        super(props);
        this.state = {
            arraypeople: [],
            arraygroups: [],
            arrayMessage: [],
            type: '',
            id_people:'',
            body:'',
            id_channel:'',
            connected: false
        }
    }


    componentWillMount() {
        this.makeListPeople();
        this.makeListGroups();

    }
    componentDidMount() {

        let ws = new WebSocket('ws://10.10.101.155:4000')
        let socket = this.socket = new Socket(ws);
        socket.on('connect', this.onConnect.bind(this));
        socket.on('disconnect', this.onDisconnect.bind(this));
        socket.on('message add', this.onMessageAdd.bind(this));
    }


    makeListPeople(){
        let token = localStorage.getItem('@websession');

        let headers = {
            Authorization: 'Bearer ' + token,
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        };

        getService(Routes.LISTPEOPLE, headers).then(data => {
            this.setState({
                arraypeople: data,
            });
        });
    }

    makeListGroups(){
        let token = localStorage.getItem('@websession');

        let headers = {
            Authorization: 'Bearer ' + token,
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        };

        getService(Routes.LISTGROUPS, headers).then(data => {
            if (data !== null){
                this.setState({
                    arraygroups: data
                })
            }

        });
    }

    findMessagePeople(e){
        let token = localStorage.getItem('@websession');

        let headers = {
             Authorization: 'Bearer ' + token,
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        };

        let body = {
            id_people: e,
        };

        postService(Routes.LISTMESSAGE, body, headers).then(data => {
            if (data.dataMessages !== null){
                this.setState({
                    arrayMessage: data.dataMessages,
                    id_people: e,
                    id_channel: data.idChannel,
                    type: 'private'
                },()=>this.setChannel());
            }else if(data.dataMessages === null && data.idChannel !== null){
                this.setState({
                    arrayMessage: [],
                    id_people: e,
                    id_channel: data.idChannel,
                    type: 'private'
                },()=>this.setChannel());
            }

        });
    }

    findMessageGroup(e){
        let token = localStorage.getItem('@websession');

        let headers = {
            Authorization: 'Bearer ' + token,
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        };

        let body = {
            id: e,
        };

        postService(Routes.LISTMESSAGEGROUP, body, headers).then(data => {
            if (data.dataMessages !== null){
                this.setState({
                    arrayMessage: data.dataMessages,
                    id_people: data.iduser,
                    id_channel: e,
                    type: 'group',
                },()=>this.setChannel());
            }else if(data.dataMessages === null && data.idChannel !== null){
                this.setState({
                    arrayMessage: [],
                    id_people: '',
                    id_channel: e,
                    type: 'group',
                },()=>this.setChannel());
            }

        });
    }

    onConnect() {
        this.setState({connected: true});
    }

    onDisconnect() {
        this.setState({connected: false});
    }



    onMessageAdd(message) {
        // this.state.arrayMessage.push(message);
        const newArrayMessage = Object.assign(this.state.arrayMessage);
        newArrayMessage.push(message);

        this.setState({
            arrayMessage: newArrayMessage
        })
    }



    // Sets the channel the user wants to talk to
    setChannel() {
        this.socket.emit('message subscribe', {channelId: this.state.id_channel});
        $("#msg_history").animate({ scrollTop: $('#msg_history').prop("scrollHeight")}, 10);
    }

    MakeMessage(){
        if (this.state.body !== ''){
            let token = localStorage.getItem('@websession');

            let headers = {
                Authorization: 'Bearer ' + token,
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            };

            let body = {
                message: this.state.body,
                id_people_message: this.state.id_people,
                id_channel: this.state.id_channel,
            };

            postService(Routes.MAKEMESSAGE, body, headers).then(data => {

                this.setState({
                    body:''
                })
            })
            $("#msg_history").animate({ scrollTop: $('#msg_history').prop("scrollHeight")}, 10);
        }

    }

    handleKeyPress = (event) => {
        if(event.key === 'Enter'){
            this.MakeMessage()
        }
    };

    render() {


        return (
                <div className="inbox_msg">
                    <div className="inbox_people">
                        <User/>
                        <div className="inbox_chat">
                            {this.state.arraypeople.map((element, index) => {
                                return(
                                    <div className={this.state.id_people === element.id ? ("chat_list active_chat"):("chat_list")} key={index} onClick={()=>this.findMessagePeople(element.id)}>
                                        <div className="chat_people">
                                            <div className="chat_img"><img
                                                src="https://ptetutorials.com/images/user-profile.png" alt="sunil"/></div>

                                            <div className="chat_ib">
                                                <h5>{element.name} {element.last_name}</h5>
                                            </div>
                                            <div className="circle circle_active"/>
                                        </div>
                                    </div>
                                )
                            })}
                            {this.state.arraygroups.length !== null && this.state.arraygroups.length > 0 && this.state.arraygroups.map((element, index) => {
                                return(
                                    <div className={this.state.id_channel === element.id ? ("chat_list active_chat"):("chat_list")} key={index} onClick={()=>this.findMessageGroup(element.id)}>
                                        <div className="chat_people">
                                            <div className="chat_img"><img
                                                src="https://ptetutorials.com/images/user-profile.png" alt="sunil"/></div>
                                            <div className="chat_ib">
                                                <h5>{element.name}</h5>
                                            </div>
                                        </div>
                                    </div>
                                )
                            })}
                        </div>
                    </div>
                    { this.state.id_channel !== '' ?
                    (<div className="mesgs">
                        <div id="msg_history" className="msg_history">
                            {this.state.arrayMessage.map((element, index) => {
                                return(<Message
                                        key={index}
                                    id_peoplemessage={element.id_people_message}
                                    id_people={this.state.id_people}
                                    type={this.state.type}
                                    createdAt={element.createdAt}
                                    author={element.author}
                                    message={element.message}
                                    />
                                    )
                            })}
                        </div>
                        <div className="type_msg">
                            <div className="input_msg_write">
                                <input type="text" className="write_msg" placeholder="Escribe Un Mensaje"
                                       value={this.state.body}
                                       onChange={(e) => this.setState({
                                           body: e.target.value
                                       })}

                                       onKeyPress={this.handleKeyPress}
                                />
                                <button onClick={()=>this.MakeMessage()} className="msg_send_btn" type="button"><i className="fa fa-paper-plane-o"
                                                                                                                   aria-hidden="true"/></button>
                            </div>
                        </div>
                    </div>):(<div/>)}
                </div>
        )
    }
}

export default ContainerChannels;
