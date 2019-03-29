import React, {Component} from 'react';
import Login from './login/log_in/index';
import Main from './main/index'
import {getService} from "./../services/services";
import {Routes} from "./../services/routes";

class HomeComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            token_auth: null,
        }
    }

    componentWillMount() {
        this.checkLocalStorage()
    }


    checkLocalStorage(){
        let token = localStorage.getItem('@websession');
        this.setState({
            token_auth: token
        });
    }

     logOut(){
         let token = localStorage.getItem('@websession');

         let headers = {
             Authorization: 'Bearer ' + token,
             'Content-Type': 'application/json',
             'Accept': 'application/json'
         };

         getService(Routes.LOGOUT, headers).then(data => {
             console.log(data)
             if (data.message === "Salio de sesion" || data.message === "invalid or expired jwt"){
                 localStorage.removeItem('@websession');
                 this.setState({
                     token_auth: null
                 });
             }
         });
     }



    render() {
        return (
          <div>
              {this.state.token_auth !== null ? <Main logout={this.logOut.bind(this)}/>: <Login CheckLocalStorage={this.checkLocalStorage.bind(this)}/>}
          </div>
        );
    }
}

export default HomeComponent;