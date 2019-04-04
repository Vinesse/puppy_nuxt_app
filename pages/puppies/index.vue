<template>
    <section class="puppies">
        <h1>{{ feedingTime() }}</h1>
        <h2>Test text. Can you see me!!!?</h2>
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
import {format} from 'date-fns'

export default {
    components: {
        Puppy
    },
    methods: {
        feedingTime() {
        return format(new Date(2014, 0, 24), 'MM/DD/YYYY');
        },
    },
    asyncData() {
        return new Promise((resolve, reject) => {
            setTimeout(() => {
                //here we are setting a timeout to fake an async request
                //that would normally be made to the server
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