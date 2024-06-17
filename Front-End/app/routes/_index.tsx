import type { MetaFunction } from "@remix-run/node";
import "~/css/dashboard.css"

export const meta: MetaFunction = () => {
  return [
    { title: "Home" },
    { name: "This is the homepage of the bravus the next level small business dashboard", content: "Homepage" },
  ];
};

export default function Index() {
  return (
    <div>
      <nav className="navBar">
        <h1 className="text-6xl isolation-auto">Bravus</h1>
      </nav>
    </div>
  );
}
