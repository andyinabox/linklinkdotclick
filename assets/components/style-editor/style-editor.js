import { debounce } from '../../lib/utils'
import { Component } from '../component'
export class StyleEditor extends Component {
  constructor() {
    super()
  }
  handleInput(evt) {
    console.log('input', evt)
  }
  handleKeyPress(evt) {
    if (evt.keyCode === 9) {
      // TODO: allow tab
    }
  }
  connectedCallback() {
    this.listen(this.slots.editor, 'input', debounce(this.handleInput))
    this.listen(this.slots.editor, 'keydown', this.handleKeyPress)
  }
}
customElements.define('linky-style-editor', StyleEditor)
