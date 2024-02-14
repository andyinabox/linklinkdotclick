import { eventsMixin } from '../lib/mixins'

export class ProgressiveForm extends HTMLFormElement {
  constructor() {
    super()
    this.onsubmit = (evt) => {
      evt.preventDefault()
      this.onSubmit(evt)
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
}
Object.assign(ProgressiveForm.prototype, eventsMixin)
