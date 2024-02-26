import { eventsMixin } from '../../lib/mixins'

export class BaseForm extends HTMLFormElement {
  static overrideSubmit = true
  constructor() {
    super()
    this.querySelectorAll('[data-hide]').forEach((el) => (el.hidden = true))

    if (this.constructor.overrideSubmit) {
      this.onsubmit = (evt) => {
        evt.preventDefault()
        this.onSubmit(evt)
      }
    }

    this.inputs = {}
    this.querySelectorAll('input[name]').forEach((el) => {
      this.inputs[el.getAttribute('name')] = el
    })
    this.buttons = {}
    this.querySelectorAll('button[name]').forEach((el) => {
      this.buttons[el.getAttribute('name')] = el
    })
  }
  onSubmit(evt) {}
  connectedCallback() {
    this.querySelectorAll('button[data-event]').forEach((el) => {
      this.listen(el, 'click', () => {
        this.broadcast(el.getAttribute('data-event'), this.state)
      })
    })
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(BaseForm.prototype, eventsMixin)
customElements.define('base-form', BaseForm, { extends: 'form' })
