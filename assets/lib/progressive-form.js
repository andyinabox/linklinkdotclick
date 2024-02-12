export class ProgressiveForm extends HTMLFormElement {
  constructor() {
    super()
    this.onsubmit = (evt) => evt.preventDefault()
    this.querySelector('input[type="text"]').remove()
  }
}
