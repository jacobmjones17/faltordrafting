<template>
  <section id="services">
    <div class="container">
      <div class="section-heading">
        <div class="section-eyebrow">Services</div>
        <h2>Drafting services for real projects and real budgets.</h2>
        <p>
          From a simple porch addition to a full custom home, FALTOR Consulting provides detailed, clear drawings
          that help your project move smoothly from idea to permit to construction.
        </p>
      </div>

      <div class="services-grid">
        <article v-for="service in services" :key="service.id" class="service-card">
          <span class="service-tag">{{ service.tag }}</span>
          <h3>{{ service.title }}</h3>
          <p>{{ service.description }}</p>
        </article>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'ServicesSection',
  data() {
    return {
      services: [
        {
          id: 1,
          tag: 'Custom Homes',
          title: 'Custom House Plans',
          description: 'Full floor plans, exterior elevations, roof plans, and basic structural layouts tailored to your lot, lifestyle, and builder.'
        },
        {
          id: 2,
          tag: 'Remodels',
          title: 'Interior Remodel Drafting',
          description: 'Open up walls, re-work kitchens, or reconfigure layouts with accurate before/after plans your contractor can follow.'
        },
        {
          id: 3,
          tag: 'Additions',
          title: 'Additions & Porches',
          description: 'Tie new square footage into existing homes — including foundations, roof connections, and exterior changes.'
        },
        {
          id: 4,
          tag: 'Garages',
          title: 'Garages & Accessory Structures',
          description: 'Detached garages, workshops, and accessory buildings designed to match your home and meet local requirements.'
        },
        {
          id: 5,
          tag: 'As-Builts',
          title: 'As-Built Drawings',
          description: 'Document existing conditions to support remodels, permits, or future planning when original plans aren\'t available.'
        },
        {
          id: 6,
          tag: 'Site Plans',
          title: 'Site & Plot Plans',
          description: 'Basic site layouts showing building footprints, driveways, setbacks, and other key information for permitting.'
        },
        {
          id: 7,
          tag: 'Permit Sets',
          title: 'Permit-Ready Construction Drawings',
          description: 'Coordinated plan sets with notes and details formatted for submittal to your municipality or builder.'
        },
        {
          id: 8,
          tag: '3D Views',
          title: '3D Renderings & Exterior Views',
          description: 'Realistic 3D perspectives that help you visualize materials, massing, and curb appeal before you build.'
        },
        {
          id: 9,
          tag: 'Other Drafting',
          title: 'Other Drafting Services',
          description: 'Need something more specific? Reach out with your project details — many custom drafting needs can be accommodated.'
        }
      ]
    }
  }
}
</script>