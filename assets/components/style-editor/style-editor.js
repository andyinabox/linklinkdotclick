import { debounce } from '../../lib/utils'
import { Component } from '../component'
export class StyleEditor extends HTMLFormElement {
  constructor() {
    super()
    this.registerSlotsMixin()
    this.registerListenMixin()

    this.userStyles = document.getElementById('user-styles')
    this.socket = new WebSocket(`wss://${window.location.host}/ws/style-editor`)
    this.socket.onmessage = (msg) => {}
    this.socket.onerror = (err) => {
      console.error(err)
    }
  }
  handleSocketMessage(msg) {
    this.slots.editor.value = msg.data
    this.userStyles.innerText = msg.data
  }
  handleSocketOpen() {
    console.log('socket open!!!')
  }
  handleInput(evt) {
    console.log('input', evt.target.value)
    this.socket.send(evt.target.value)
  }
  handleKeyPress(evt) {
    if (evt.keyCode === 9) {
      // TODO: allow tab
    }
  }
  connectedCallback() {
    this.listen(
      this.slots.editor,
      'input',
      debounce(this.handleInput.bind(this), 500)
    )
    this.listen(this.slots.editor, 'keydown', this.handleKeyPress)
    this.listen(this.socket, 'open', this.handleSocketOpen)
  }
}
Object.assign(Component.prototype, slotsMixin, listenMixin)

customElements.define('linky-style-editor', StyleEditor, { extends: 'form' })
