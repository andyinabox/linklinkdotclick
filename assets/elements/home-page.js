class HomePage extends HTMLBodyElement {
  constructor() {
    super()
    console.log('create home page')
  }
}
customElements.define('linky-home-page', HomePage, { extends: 'body' })
