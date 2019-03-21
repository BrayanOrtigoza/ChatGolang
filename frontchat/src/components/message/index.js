import React, { Component } from 'react';
import fecha from 'fecha';

let createdAt;
class Message extends Component {

    constructor(props) {
        super(props);
        this.state = {
            body: '',
            arrayMessage:[],
            connected: true
        }
    }

    componentWillMount() {
        createdAt = fecha.format(new Date(this.props.createdAt), 'DD/MM HH:mm:ss');
        console.log(this.props.id_peoplemessage)
        console.log(this.props.id_people)
    }


    render() {
        return (
                <div>
                    {(this.props.id_peoplemessage !== this.props.id_people && this.props.type === 'private') ? (
                        <div  className="outgoing_msg">
                            <div className="sent_msg">
                                <span className="time_date">{this.props.author}</span>
                                <p>{this.props.message}</p>
                                <span className="time_date">{createdAt}</span></div>
                        </div>
                        ):(this.props.id_peoplemessage === this.props.id_people && this.props.type === 'group')? (
                        <div  className="outgoing_msg">
                            <div className="sent_msg">
                                <span className="time_date">{this.props.author}</span>
                                <p>{this.props.message}</p>
                                <span className="time_date">{createdAt}</span></div>
                        </div>
                        ):(
                            <div className="incoming_msg">
                                <div className="incoming_msg_img">
                                    <img src="https://ptetutorials.com/images/user-profile.png" alt="sunil"/>
                                </div>
                                <div className="received_msg">
                                    <div className="received_withd_msg">
                                        <span className="time_date">{this.props.author}</span>
                                        <p>{this.props.message}</p>
                                        <span className="time_date">{createdAt}</span>
                                    </div>
                                </div>
                            </div>
                        )}
                </div>
        )
    }
}
export default Message;
