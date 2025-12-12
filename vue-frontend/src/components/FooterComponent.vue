<template>
  <footer>
    <div class="container footer-row">
      <div>© {{ currentYear }} FALTOR Consulting. All rights reserved.</div>
      <div>
        Drafting services for custom homes, remodels, additions, and more across Eastern North Carolina.
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'FooterComponent',
  computed: {
    currentYear() {
      return new Date().getFullYear()
    }
  }
}
</script>