import React, { Component } from 'react';
import '../../../styles/login.css'
import {postService} from "../../../services/services";
import {Routes} from "../../../services/routes";



class LogIn extends Component {

    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: '',
            message:'',
        }
    }

    handleValidate() {
        return this.state.username !== '' && this.state.password !== '';
    }

    handleLogin = (e) => {
        e.preventDefault();
        this.handleValidate() ? this.sendLogin() : this.showErrorMesage();
    };


    showErrorMesage(){
        this.setState({
            message: 'incorrecto'
        });
        console.log('falta usuario y/o contraseña')
    }


    sendLogin() {
        let headers = {
            'Content-Type': 'application/json'
        };

        let body = {
            username: this.state.username,
            password: this.state.password,
        };

        postService(Routes.LOGIN, body, headers).then(data => {
            if(data.token !== null && data.token !== undefined) {
                localStorage.setItem("@websession", data.token);
                this.props.CheckLocalStorage();
            }else{
                this.setState({
                    message: data.error,
                });
                console.log(data.error)
            }
        })
    }

    render() {
        return (
            <div className="wrapper fadeInDown">
                <div id="formContent">

                    <h2 className="active"> Sign In</h2>


                    <div className="fadeIn first">
                        <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANIAAADSCAMAAAAIR25wAAAAYFBMVEX///8AmtoKntz2/P4Vo97g8/smquD3/P7t+P3T7vkFnNtdv+gureHx+v3A5/b7/v+z4fSQ1O98zOxlw+lxx+um3PNOueaJ0e6e2fE+s+MsrOGs3/PK6/gcpt/a8fp1yeurAlnEAAALZklEQVR4nO1d65qrKgydsV6LWhGvHa3v/5ZHCFptrddg3X5n/Zu9W0skCSshhJ+f/7EQnlEEcUZpGoYppVkcFIb37TGtxM2JU5aY2u8ANDNhaezcvj3G+bj4KXEHZXmRzCWpf/n2aCdh/0XutTfu0s0JYVEVhlXECMndsift1Y3+7G+P+jOKMH+OtUwiGjiDduM5AY2S8vnZPCz2HuscFExvpWHUn+EDPJ+yVi6dHUwqIzUbcUhmLPpmRhqxzHTRN1XCy6S+6SR7rHnAIyNyhvPsCB7eJ2DuGvM3PYbJx5BNj0FAABN0TWJr66OsOAFfmQcYI1uJ2BVjcCmSG7apfGCM87yluGUmKJyD+VQHFNDM9ucWFhWOSqvQ10m7EkKVdLMqL0MgZkgPlSz8digcoLmnTRmJEChV5nG9VAiV7LVQWSnXDJ0qZZwXyoXS0l20zxc6x5RzTZsJ7VO/TBlEeNldCFkhXDpRrH0x1zmNqv2RJ6j4OZWrlCV0gewY39hCKZgyi3q4+yh3D8J03VV0eBpC6cjuVNkjqpTvwpVOyxQ8eRIZf5cMfc1w7lzpUOncgh/nyndH/vFAKN3X0jkXoXyoBCm+fkvpGnDluyIaFP2i0jUQyoe2IFbcjX491WHwJaRCedSNu7r8AGkOjycFGEJoeOGBRLJzODYMSwxls4+ysV4NBkBhNvKxC5coxBkQBkIu06Z5uiWYbgYD3PkmW5SGHU0ikImt/3p1LK0DhFt8+cYXogobVCfeqraKIAx8FTcKal6XH2I9eoVVu+HrCg7r1DzRPQBnGIJXcyNtMee81PGR+XVe9wlGzWHvS5cnsuY97AeuQ2TZV7hr+Gp8NIVsqYt4LH8Je4Or0YK8kTCkgxdZXJaZEzu2IQG4Oc0mAoc3JMACczKOb0gAbk7z1pn6k+ZB19g+PHPmu/fr+fx27cFMzByqNVf0I4Ar1DQNTWsFPXA1WR92bfbp1Ie4bzhYHDsGOsND1LGIu8tgkFBz8mT8E0FtcAcrjBtHUQ94NHTivuGAofkY2ISHqFVT/2d8A8DWR43/ViryDU5KXF27amW+sKhyBuppKD9nSGrWpOMTcD/qVObybGmKupl80ccYqTnDyy+F7/6+A7U8o15JzU//VzNwHZncOcmAQDWuEZ7JevpnRu6i51bTYYE4dDweGX5cSwNsKmSRzxLVE4UWknFaNLw25Vj7hhLG04queRo4tmUUf5X5FCrCSuVWtc8Z+ncfeZK8dvBl1n3uo2oPXWCt6nyahvSYIBOHxjHob2WCBmlOmmAtgmwwIPJqSTFTKJUc9eA+pC+r969IPsKpNezdV2e4FPyvWYGG+Zdxl1OIpOru0HKbo3IhSxKG6NMHPBfVnOiAgzBqLUB0DhQGPLI9ZYDuXXHIkV1b5ysjCSdDqSXwYLyjiSYffATSrybvNMFcubE2jBAmadz4I/gQTsQZvxG9OjjUEHf8wPgHl78nbJhKnOXd0l5fDkNdlB7w/qfWBJjLO85vvkmgo6Yjga1OrgkGSI7jIGryo3f/Ll7+3gg4pzUderkzPzcLel/zQtQM6+06S++a2UTyeaTv83LU3RcwpRnT7mMaU9bzRzaaRgsEc0cKsms4v8of9iQLNSErcZ7bPG7ahXNcwD8ghU01B/tr/4hwNysyIKwzPgkrExIRI11KOchj1yOebfXgR5DybN1Y4jLA+bYgmKt4Nige0s9y5t28HR/XlETqfdYji7mucR7KJ19IkUN0ePnX6ZRgPHc654E9l22CnQkHq/+b/BzBjAJFkNb4JBd7v5nNc3k3EB3tdIjf+odbzctx88bgH/SpaAU+hhfUePXDYI1zRtLk62Bps/hojknxOMqGWMaojxWAHN5E+kcmkRBXxKQJzdORRM5KAB8df6wF2VjMDa2o0QymYOtPplrHnB5BnyTu8ljz8+hn9h0Yr/Y5VyL3aUzMqu2gMSETN3MMAD/+W356MlDbOYvXArSODt2Hc9hy30IbHPNNJryQq5U8GXx5aEFYD07Taal6f1+PpqsUdmm9nB0DmbQ2COSwf/WXliEGa/ZiSuwaixJCikJV1RD9bYViQSOVnSXtjpmOXtrjQpYoQOTCfQSdHmvaPSERyTvbmr93/MMCOTjvWF1JodOV4BWJgmJTAvQhU1gJZX/cU9fRa0Y4GCzdVGlxV5EPCYRZydGFJEIUn+L1EOTX1xkiijpScJLHRUqRyx3e4cXk6SjMyFd3cq0C3hrucoTRCDKa0thXW+0nZdlHpH0gZVGveBKW+iMCUvFUuwcjSFl+L3XRUkg33aSKHVWnJqPW4ylz4g+aD/akvbpq2ppKJ56pYg+Paow8/JYK2poSWGoVEaLiQ71kFzn2cRVJiJTQ1mHSoFwoSVsVBBdW9UoYhAHpQ/8aYZY8y+ACPwQsejak54wGjmFzznCxH0VWJb3/x2wSJkNA9EA97kzGh4bNPU+IV+TaplGQ0ymdYLYaCfLstG0DjVac3s4ObtIrbAWaakZpZW2VPxJ9aZNeqKnJdo7mNAa0WfNpnHRvm5rETCAX0o70mS8pQC1ybYkQYprflqr0Mcn6hof8ho6RW2nT/IibMZIxLGkUYUiPjrHct5sxDtrCFKx547LIFWMLo90yQ9vYvMlK6YXORha5bj+U89zYRNt+lt5ucYxMkTz5c/sZq0hAbpCXy+M7KDXcXNTdKRJAKuWQlrRix8hfOb0v6JRyIBXcJOs9FyRlJ6sKJtApuMEpi7LByldVmxerJ7iDblkUTvEaGPnKDXLwldui695BGJQSw2TLqOB9bCtU6ZUYohSCgr9buRrY2pZvA3qFoBjlulAWoK3NdyebfV6/XBejqBpUZ3UWI91M9PpF1Ril77Dtv3p9g6VpSwHlS+k7wgGFZBv39MCNb1hLXg4oIBwjcTfa930N4+3g9RgJwmEfcHjrnQzZGGG8SbD9SBZ44fUBAtji6lzR+5Gs7QfnZtyaNwOrA4z3g3PbjzfiiLRa+5P3Gd58CBVHpLVLydAh1M1Hhcvp8U7KYzzWOvGho8KbD3QP9RPZaYp+PsQSW4/dz9ghUyfS8LF79OYIe2K4OQJ6C4sd8amFBXqjkf3wqdEIfjuYvfC5HYyCpj374HPTHhWtlfbAWGslJQ2w1GOsAZaiNmWKMd6mTFkzOZUYbyZ3wpZ/Z2zMeML2mWdscnrGVrQnbBh8xrbOJ2y+fcYW6WdsZH/C6wbOeCnEj+X+G1d3uAvS+Oe7YOUfMKfF1+Ac3pyWGRLghFdKnfDirzNez3bGS/TOeNXhGS+kPOO1oWe83PWEV/Ce8aLkM15nfcZLx39OeDX8D7gZ88u8XDQtQXS+/MSi9tX4KatZ9BWx47lsVkO+FrtfeLWehtx5zLl/UfmE0t3Rfxxe1FeUL1OmIrF48u7e3BPvEtWMnngI5ds5X+4LpVPVEufnIs7BzjkxiwXozcRU+iWhfNpu3JxqCpWugSFem7vLfm4hyhZR7z0bhlDuX6Zc++C09z6ma6VcHSZP1W/DhfJaeu3t7iZVMETVp54qc+gedE5I9uTKgdA+PVSifnYoBDLRW6+Ow6Kiilqr0IWy4a6zku4fdt4yMVMaQ+VeDhMCmdmXcgMxFIe7FGmqbCofqHglGkUA/YauSbxZTaw4gZO5+c429AZftvPT2KYVxGfyMeQINRdeJltD6SRbRS8fGZE9H/LsAGkbgBE23ZLKhXfsGhlpDqCY4ddTNn0UrO0rVDLqz3jbnk9Ze55GV9CHDQFF2GlOZiYRDZxByTwnoFG3C9aHnlLHgP0Xub1mZFrp5oSxqArDKmKM5G7ZO7x1daO/49fKXfyUuDMOnWkuSf2DF4x0cXPilCXmoGSambA0dg6yd7AYnlEEcUZpGoYppVkcFMZh/PS/g/8AmEt4+y03XTwAAAAASUVORK5CYII=" id="icon" alt="User Icon"/>
                    </div>


                    <form onSubmit={this.handleLogin}>
                        <input type="text" className="fadeIn second"
                               placeholder="Usuario"
                               value={this.state.username}
                               onChange={(e) => this.setState({
                                   username: e.target.value
                               })}
                        />
                        <input type="password" id="password" className="fadeIn third"
                               placeholder="contraseña"
                               value={this.state.password}
                               onChange={(e) => this.setState({
                                   password: e.target.value
                               })}
                        />
                        <input type="submit" className="fadeIn fourth" value="Iniciar sesion"/>
                    </form>
                </div>
            </div>
        );
    }
}

export default LogIn;
