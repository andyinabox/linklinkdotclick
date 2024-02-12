export class ProgressiveForm extends HTMLFormElement {
  constructor() {
    super()
    this.onsubmit = (evt) => evt.preventDefault()
  }
}
