export class Component extends HTMLElement {
  constructor() {
    super()
    this.slots = {}
    this.querySelectorAll('[slot]').forEach((el) => {
      this.slots[el.getAttribute('slot')] = el
    })
  }
}
