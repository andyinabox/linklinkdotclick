import { debounce } from '../../lib/utils'
import { handleError } from '../../lib/errors'
import { FormBase } from './form-base'
export class FormStyleEditor extends FormBase {
  static overrideSubmit = false
  constructor() {
    super()
    this.editor = this.querySelector('textarea')
    this.errors = this.querySelector('ul.errors')
    this.originalStyles = this.editor.value
    this.socket = new WebSocket(`wss://${window.location.host}/ws/style-editor`)
    this.setButtonStatus()
  }

  set validCss(bool) {
    if (bool) {
      this.classList.remove('invalid')
    } else {
      this.classList.add('invalid')
    }
  }
  get validCss() {
    return !this.classList.contains('invalid')
  }

  get prevChar() {
    return this.editor.value[this.editor.selectionStart - 1]
  }

  set userStyles(styles) {
    document.getElementById('user-styles').innerHTML = `
  @layer user {
    ${styles}
  }`
  }

  setButtonStatus() {
    if (this.editor.value === this.originalStyles) {
      this.buttons['btn-styles-reset'].disabled = true
      this.buttons['btn-styles-save'].disabled = true
    } else {
      this.buttons['btn-styles-reset'].disabled = false
      this.buttons['btn-styles-save'].disabled = false
    }
    if (!this.validCss) {
      this.buttons['btn-styles-save'].disabled = true
    }
  }

  sendData(styles) {
    try {
      const data = { styles }
      console.log('websocket request', data)
      const jsonData = JSON.stringify(data)
      this.socket.send(jsonData)
    } catch (err) {
      handleError(err)
    }
  }

  handleSocketMessage(msg) {
    try {
      const data = JSON.parse(msg.data)
      console.log('websocket response', data)
      this.validCss = data.valid
      this.renderErrors(data.warnings)
      if (data.valid) {
        const selectionStart = this.editor.selectionStart
        const selectionEnd = this.editor.selectionEnd
        this.editor.value = data.styles
        this.userStyles = data.styles
        this.editor.setSelectionRange(selectionStart, selectionEnd)
        this.setButtonStatus()
      }
    } catch (err) {
      handleError(err)
    }
  }

  handleInput(evt) {
    if (this.prevChar == ' ') return
    this.sendData(this.editor.value)
  }

  onReset() {
    window.requestAnimationFrame(() => this.sendData(this.editor.value))
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

  renderErrors(errors) {
    this.errors.innerHTML = ''

    if (!errors.length) {
      this.errors.hidden = true
      return
    }
    errors.forEach((e) => {
      const el = document.createElement('li')
      el.innerText = e
      this.errors.appendChild(el)
    })
    this.errors.hidden = false
  }

  handleKeyPress(evt) {
    // tab key
    if (evt.keyCode === 9) {
      evt.preventDefault()
      this.insert('  ')
    }
    // return key
    if (evt.keyCode === 13) {
      console.log('return key')
      if (this.prevChar === '{' || this.prevChar === ';') {
        evt.preventDefault()
        this.insert('\n  ')
      }
    }
  }

  connectedCallback() {
    super.connectedCallback()
    this.listen(
      this.editor,
      'input',
      debounce(this.handleInput.bind(this), 500)
    )
    this.listen(this, 'reset', this.onReset)
    this.listen(this.editor, 'keydown', this.handleKeyPress)
    this.listen(this.socket, 'message', this.handleSocketMessage)
  }
}

customElements.define('form-style-editor', FormStyleEditor, { extends: 'form' })
