import React, { Component } from 'react';
import GroupsList from './groupsList'

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
                    <GroupsList />
                    <ChannelForm {...this.props} />
                </div>
            </div>
        )
    }
}

export default Main;
