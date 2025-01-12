<script lang="ts">
    import "../app.css";
    let { children } = $props();
</script>

<article class="bg-gray-50 dark:bg-zinc-900 text-black dark:text-white w-[100dvw] h-[100dvh] py-4 px-8">
    {@render children()}
</article>
