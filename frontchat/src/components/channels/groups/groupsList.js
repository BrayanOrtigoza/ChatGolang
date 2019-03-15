import React, { Component } from 'react';

class Main extends Component {

    componentWillMount() {

    }



    render() {
        return (
            <div className='support panel panel-primary'>
                <div className='panel-heading'>
                    <strong>Channels</strong>
                </div>
                <div className='panel-body channels'>
                    <ChannelList />
                    <ChannelForm />
                </div>
            </div>
        )
    }
}

export default Main;
