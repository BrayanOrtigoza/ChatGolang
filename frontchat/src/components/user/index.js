import React, { Component } from 'react';
import {getService} from "../../services/services";
import {Routes} from "../../services/routes";



class Main extends Component {

    constructor(props) {
        super(props);
        this.state = {
            id:'',
            name: '',
            last_name: '',
        }
    }

    componentWillMount() {
        this.findDataUser()
    }


    findDataUser(){
        let token = localStorage.getItem('@websession');

        let headers = {
            Authorization: 'Bearer ' + token,
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        };

        getService(Routes.FINDUSER, headers).then(data => {
            if(data.error !== undefined){
                localStorage.removeItem('@websession');
                setTimeout(function () {
                    window.location.href = "/";
                }, 2000);
            }else{
                this.setState({
                    name: data.name,
                    last_name: data.last_name,
                });
            }
        })
    }

    render() {
        return (
                <div className="headind_srch">
                    <div className="recent_heading">
                        <h4>{this.state.name} {this.state.last_name}</h4>
                    </div>

                </div>
        )
    }
}

export default Main;
