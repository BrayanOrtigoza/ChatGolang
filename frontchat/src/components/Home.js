import React, {Component} from 'react';
import Login from './login/log_in/index';
import Main from './main/index'

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



    render() {
        return (
          <div>
              {this.state.token_auth !== null ? <Main/>: <Login CheckLocalStorage={this.checkLocalStorage.bind(this)}/>}
          </div>
        );
    }
}

export default HomeComponent;