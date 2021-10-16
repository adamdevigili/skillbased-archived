import Link from "next/link";
import useSWR from "swr";

async function fetcher(url: string) {
  const resp = await fetch(url);
  return resp.text();
}

function Index(): JSX.Element {
  const { data, error } = useSWR("/v1/health", fetcher);

  return (
    <div>
      <h1>Skillbased</h1>

      <p>
        Check out our <Link href="/api">API</Link>.
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
