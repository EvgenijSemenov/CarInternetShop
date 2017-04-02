export class Authorization {

  public isAuthorizedUser(): boolean {
    let authToken = localStorage.getItem('auth-token');
    if (authToken != null) {
      return true;
    }
    return false;
  }

  public authorizateUser() {
    localStorage.setItem('auth-token', "token");
  }

  public removeToken() {
    localStorage.removeItem('auth-token');
  }

}
