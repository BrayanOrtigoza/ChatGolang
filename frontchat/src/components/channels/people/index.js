import React, { Component } from 'react';

class People extends Component {


    render() {
        return (
                <div className="inbox_chat">
                    {this.props.Arraypeople.map((element, index) => {
                    return(
                        <div className="chat_list active_chat" key={index}>
                            <div className="chat_people">
                                <div className="chat_img"><img
                                    src="https://ptetutorials.com/images/user-profile.png" alt="sunil"/></div>
                                <div className="chat_ib">
                                    <h5>{element.name} {element.last_name}</h5>
                                </div>
                            </div>
                        </div>
                      )
                    })}
                </div>
        )
    }
}

export default People;
