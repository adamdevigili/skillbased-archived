import Link from "next/link";
import useSWR from "swr";

// async function fetcher(url: string) {
//   const resp = await fetch(url);
//   return resp.text();
// }

const fetcher = (url) => fetch(url).then((res) => res.json());

function Index(): JSX.Element {
  const { data, error } = useSWR("/v1/health", fetcher, { refreshInterval: 1000 });

  console.log(data)

  return (
    <div>
      <h1>Hello, world!</h1>
      <p>
        This is <code>pages/index.tsx</code>.
      </p>
      <p>
        Check out <Link href="/foo">foo</Link>.
      </p>

      <h2>Memory allocation stats from Go server</h2>
      {error && (
        <p>
          Error fetching profile: <strong>{error}</strong>
        </p>
      )}
      {!error && !data && <p>Loading ...</p>}
      {!error && data && <pre>{data}</pre>}
    </div>
  );
}

export default Index;
