import { debounce } from '../../lib/utils'
import { FormBase } from './form-base'
export class FormStyleEditor extends FormBase {
  static overrideSubmit = false
  constructor() {
    super()
    this.editor = this.querySelector('textarea')
    this.userStyles = document.getElementById('user-styles')
    // this.buttons['btn-styles-open'].hidden = false
    this.socket = new WebSocket(`wss://${window.location.host}/ws/style-editor`)
    this.socket.onerror = (err) => {
      console.error(err)
    }
  }
  // handleNewWindowClick() {
  //   this.popup = window.open(
  //     '/styles',
  //     '_blank',
  //     'popup=1,width=500,height=700'
  //   )
  // }

  handleSocketMessage(msg) {
    const selectionStart = this.editor.selectionStart
    const selectionEnd = this.editor.selectionEnd
    this.editor.value = msg.data
    this.userStyles.innerText = msg.data
    this.editor.setSelectionRange(selectionStart, selectionEnd)
  }
  // handleSocketOpen() {
  //   console.log('socket open')
  // }
  handleInput(evt) {
    if (this.prevChar == ' ') return
    this.socket.send(this.editor.value)
  }

  onReset() {
    window.requestAnimationFrame(() => this.socket.send(this.editor.value))
  }

  get prevChar() {
    return this.editor.value[this.editor.selectionStart - 1]
  }

  insert(str) {
    const selectionStart = this.editor.selectionStart
    const selectionEnd = this.editor.selectionEnd
    const start = this.editor.value.slice(0, selectionStart)
    const end = this.editor.value.slice(selectionStart)
    this.editor.value = start + str + end
    this.editor.setSelectionRange(
      selectionStart + str.length,
      selectionEnd + str.length
    )
  }

  handleKeyPress(evt) {
    // tab key
    if (evt.keyCode === 9) {
      evt.preventDefault()
      this.insert('  ')
    }
    if (evt.keyCode === 13) {
      console.log('return key')
      if (this.prevChar === '{' || this.prevChar === ';') {
        evt.preventDefault()
        this.insert('\n  ')
      }
    }
  }
  connectedCallback() {
    this.listen(
      this.editor,
      'input',
      debounce(this.handleInput.bind(this), 500)
    )
    this.listen(this, 'reset', this.onReset)
    this.listen(this.editor, 'keydown', this.handleKeyPress)
    // this.listen(this.socket, 'open', this.handleSocketOpen)
    this.listen(this.socket, 'message', this.handleSocketMessage)
    // this.listen(
    //   this.buttons['btn-styles-open'],
    //   'click',
    //   this.handleNewWindowClick
    // )
  }
}

customElements.define('form-style-editor', FormStyleEditor, { extends: 'form' })
