<template>
    <section class="puppies">
        <Puppy
            v-for="puppy in puppies"
            :key="puppy.id"
            :thumbnail="puppy.thumbnail"
            :title="puppy.title"
            :id="puppy.id"
            :previewText="puppy.previewText" />
    </section>
</template>

<script>
import Puppy from '@/components/Puppy'

export default {
    components: {
        Puppy
    },
    asyncData() {
        return new Promise((resolve, reject) => {
            setTimeout(() => {
                //we are faking an async request/faking fetching data from a server
                resolve({
                    puppies: [
                        {
                            id: "1",
                            title: "Pug",
                            thumbnail: "http://cdn-www.dailypuppy.com/media/dogs/anonymous/8684/2008091699067_100_2824.jpg_w450.jpg",
                            previewText: "Really cute Pug",
                        },
                        {
                            id: "2",
                            title: "Boston Terrier",
                            thumbnail: "http://www.puppy-4-sale.net/black-boston-terrier-puppy.jpg",
                            previewText: "Really cute Boston Terrier",
                        }
                    ]
                })
            }, 1500)
        })
    }
}
</script>


<style scoped>
.puppies {
    display: flex;
    flex-flow: row wrap;
    justify-content: center;
    align-items: center;
}
</style>