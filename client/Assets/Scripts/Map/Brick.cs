using UnityEngine;
using System.Collections;
using System.Collections.Generic;

[RequireComponent(typeof(MeshRenderer))]
[RequireComponent(typeof(MeshCollider))]
[RequireComponent(typeof(MeshFilter))]
public class Brick : MonoBehaviour
{


    public Mesh visualMesh;
    protected MeshRenderer meshRenderer;
    protected MeshCollider meshCollider;
    protected MeshFilter meshFilter;

    // Use this for initialization
    void Start()
    {


        meshRenderer = GetComponent<MeshRenderer>();
        meshCollider = GetComponent<MeshCollider>();
        meshFilter = GetComponent<MeshFilter>();

        StartCoroutine(CreateVisualMesh());

    }

    // Update is called once per frame
    void Update()
    {

    }


    public virtual IEnumerator CreateVisualMesh()
    {
        visualMesh = new Mesh();

        List<Vector3> verts = new List<Vector3>();
        List<Vector2> uvs = new List<Vector2>();
        List<int> tris = new List<int>();
        int x = 0, y = 0, z = 0;

        byte brick = 125;
        // Left wall
        BuildFace(brick, new Vector3(x, y, z), Vector3.up, Vector3.forward, false, verts, uvs, tris);
        // Right wall
        BuildFace(brick, new Vector3(x + 1, y, z), Vector3.up, Vector3.forward, true, verts, uvs, tris);

        // Bottom wall
        BuildFace(brick, new Vector3(x, y, z), Vector3.forward, Vector3.right, false, verts, uvs, tris);
        // Top wall
        BuildFace(brick, new Vector3(x, y + 1, z), Vector3.forward, Vector3.right, true, verts, uvs, tris);

        // Back
        BuildFace(brick, new Vector3(x, y, z), Vector3.up, Vector3.right, true, verts, uvs, tris);
        // Front
        BuildFace(brick, new Vector3(x, y, z + 1), Vector3.up, Vector3.right, false, verts, uvs, tris);

        visualMesh.vertices = verts.ToArray();
        visualMesh.uv = uvs.ToArray();
        visualMesh.triangles = tris.ToArray();
        visualMesh.RecalculateBounds();
        visualMesh.RecalculateNormals();

        meshFilter.mesh = visualMesh;

        meshCollider.sharedMesh = null;
        meshCollider.sharedMesh = visualMesh;

        yield return 0;

    }
    public virtual void BuildFace(byte brick, Vector3 corner, Vector3 up, Vector3 right, bool reversed, List<Vector3> verts, List<Vector2> uvs, List<int> tris)
    {
        int index = verts.Count;

        verts.Add(corner);
        verts.Add(corner + up);
        verts.Add(corner + up + right);
        verts.Add(corner + right);

        Vector2 uvWidth = new Vector2(0.25f, 0.25f);
        Vector2 uvCorner = new Vector2(0.00f, 0.75f);

        uvCorner.x += (float)(brick - 1) / 4;

        uvs.Add(uvCorner);
        uvs.Add(new Vector2(uvCorner.x, uvCorner.y + uvWidth.y));
        uvs.Add(new Vector2(uvCorner.x + uvWidth.x, uvCorner.y + uvWidth.y));
        uvs.Add(new Vector2(uvCorner.x + uvWidth.x, uvCorner.y));

        if (reversed)
        {
            tris.Add(index + 0);
            tris.Add(index + 1);
            tris.Add(index + 2);
            tris.Add(index + 2);
            tris.Add(index + 3);
            tris.Add(index + 0);
        }
        else
        {
            tris.Add(index + 1);
            tris.Add(index + 0);
            tris.Add(index + 2);
            tris.Add(index + 3);
            tris.Add(index + 2);
            tris.Add(index + 0);
        }

    }

}


