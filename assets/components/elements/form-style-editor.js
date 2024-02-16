import { debounce } from '../../lib/utils'
import { FormBase } from './form-base'
export class FormStyleEditor extends FormBase {
  constructor() {
    super()
    this.editor = this.querySelector('textarea')
    this.userStyles = document.getElementById('user-styles')
    this.buttons['btn-styles-open'].hidden = false
    this.socket = new WebSocket(`wss://${window.location.host}/ws/style-editor`)
    this.socket.onerror = (err) => {
      console.error(err)
    }
  }
  handleNewWindowClick() {
    this.popup = window.open(
      '/styles',
      '_blank',
      'popup=1,width=500,height=700'
    )
  }

  handleSocketMessage(msg) {
    console.log('socket message', msg)
    this.editor.value = msg.data
    this.userStyles.innerText = msg.data
  }
  handleSocketOpen() {}
  handleInput(evt) {
    this.socket.send(evt.target.value)
  }
  handleKeyPress(evt) {
    if (evt.keyCode === 9) {
      // TODO: allow tab
    }
  }
  connectedCallback() {
    this.listen(
      this.editor,
      'input',
      debounce(this.handleInput.bind(this), 500)
    )
    this.listen(this.editor, 'keydown', this.handleKeyPress)
    this.listen(this.socket, 'open', this.handleSocketOpen)
    this.listen(this.socket, 'message', this.handleSocketMessage)
    this.listen(this.socket, 'message', this.handleSocketMessage)
    this.listen(
      this.buttons['btn-styles-open'],
      'click',
      this.handleNewWindowClick
    )
  }
}

customElements.define('form-style-editor', FormStyleEditor, { extends: 'form' })
