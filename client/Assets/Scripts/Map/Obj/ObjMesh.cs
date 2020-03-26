using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace NMObj
{
    public partial class Obj
    {

        private List<int> m_Indices = new List<int>();
        private List<Vector2> m_Uvs = new List<Vector2>();
        private List<Vector3> m_Vertices = new List<Vector3>();
        private List<Color> m_Colors = new List<Color>();
        private static int m_Id;
        private MeshFilter m_meshFilter;
        private MeshCollider m_meshCollider;
        private MeshRenderer m_meshRenderer;
        private int m_meshX = Config.TerrainMeshX;
        private int m_meshY = Config.TerrainMeshY;
        private int m_meshZ = 0;

        private Vector3 m_size = new Vector3(10,10,10);

        public Vector3 Size {
            get { return m_size; }
            set { m_size = value; }
        }

        public int MeshX
        {
            get { return m_meshX; }
            set { m_meshX = value; }
        }

        public int MeshY
        {
            get { return m_meshY; }
            set { m_meshY = value; }
        }

        public int MeshZ
        {
            get { return m_meshZ; }
            set { m_meshZ = value; }
        }

        public List<int> Indices
        {
            get { return m_Indices; }
            set { m_Indices = value; }
        }

        public List<Vector2> Uvs
        {
            get { return m_Uvs; }
            set { m_Uvs = value; }
        }

        public List<Vector3> Vertices
        {
            get { return m_Vertices; }
            set { m_Vertices = value; }
        }

        public List<Color> Colors
        {
            get { return m_Colors; }
            set { m_Colors = value; }
        }


        public void InitMeshData()
        {

            m_meshRenderer = ObjTransform.GetComponent<MeshRenderer>();
            m_meshCollider = ObjTransform.GetComponent<MeshCollider>();
            m_meshFilter = ObjTransform.GetComponent<MeshFilter>();
            ObjTransform.localScale = InitScale();
            ObjTransform.position += new Vector3(0, OffY(), 0);
            //CalculateMapFromScratch();
            CreateVisualMesh();


        }

        //public void Light()
        //{
        //    m_meshFilter.mesh.Clear();
        //    m_meshFilter.mesh.vertices = Vertices.ToArray();
        //    m_meshFilter.mesh.uv = Uvs.ToArray();
        //    m_meshFilter.mesh.colors = Colors.ToArray();
        //    m_meshFilter.mesh.triangles = Indices.ToArray();

        //    m_meshFilter.mesh.RecalculateBounds();
        //    m_meshFilter.mesh.RecalculateNormals();

        //    m_meshCollider.sharedMesh = null;
        //    m_meshCollider.sharedMesh = m_meshFilter.mesh;

        //    Vertices = new List<Vector3>();
        //    Uvs = new List<Vector2>();
        //    Colors = new List<Color>();
        //    Indices = new List<int>();
        //}

        public virtual void GenerateMeshData()
        {
            for (int x = 0; x < MeshX; x++)
            {
                for (int z = 0; z < MeshY; z++)
                {

                    int y = 0;
                    byte brick = 1;
                    // Left wall
                    // BuildFace(brick, new Vector3(x, y, z), Vector3.up, Vector3.forward, false);
                    // Right wall
                    // BuildFace(brick, new Vector3(x + 1, y, z), Vector3.up, Vector3.forward,true);

                    // Bottom wall
                    // BuildFace(brick, new Vector3(x, y, z), Vector3.forward, Vector3.right, false);
                    // Top wall
                    BuildFace(brick, new Vector3(x, y + 1, z), Vector3.forward, Vector3.right, true);

                    // Back
                    // BuildFace(brick, new Vector3(x, y, z), Vector3.up, Vector3.right, true);
                    // Front
                    // BuildFace(brick, new Vector3(x, y, z + 1), Vector3.up, Vector3.right, false);


                }
            }
        }


        public void CreateVisualMesh()
        {
            var visualMesh = new Mesh();
            Vertices = new List<Vector3>();
            Uvs = new List<Vector2>();
            Indices = new List<int>();

            GenerateMeshData();

            visualMesh.vertices = Vertices.ToArray();
            visualMesh.uv = Uvs.ToArray();
            visualMesh.triangles = Indices.ToArray();
            visualMesh.RecalculateBounds();
            visualMesh.RecalculateNormals();

            m_meshFilter.mesh = visualMesh;

            m_meshCollider.sharedMesh = null;
            m_meshCollider.sharedMesh = visualMesh;


        }
        public virtual void BuildFace(byte brick, Vector3 corner, Vector3 up, Vector3 right, bool reversed)
        {
            int index = Vertices.Count;

            Vertices.Add(corner);
            Vertices.Add(corner + up);
            Vertices.Add(corner + up + right);
            Vertices.Add(corner + right);




            SetTextureAtlasUv();



            if (reversed)
            {
                Indices.Add(index + 0);
                Indices.Add(index + 1);
                Indices.Add(index + 2);
                Indices.Add(index + 2);
                Indices.Add(index + 3);
                Indices.Add(index + 0);
            }
            else
            {
                Indices.Add(index + 1);
                Indices.Add(index + 0);
                Indices.Add(index + 2);
                Indices.Add(index + 3);
                Indices.Add(index + 2);
                Indices.Add(index + 0);
            }

        }

        public virtual Vector3 InitScale()
        {
            return Config.Scale * Size;
        }

        /// <summary>
        /// y轴偏移量
        /// </summary>
        /// <returns></returns>
        public virtual int OffY()
        {
            return 0;
        }

        public virtual void SetTextureAtlasUv()
        {
            const float epsilon = 0.001f;
            var TextureAtlasUv = GetTextureAtlasUv();

            Uvs.Add(new Vector2(TextureAtlasUv.x + epsilon, TextureAtlasUv.y + epsilon));
            Uvs.Add(new Vector2(TextureAtlasUv.x + epsilon,
                                      TextureAtlasUv.y + TextureAtlasUv.height - epsilon));
            Uvs.Add(new Vector2(TextureAtlasUv.x + TextureAtlasUv.width - epsilon,
                                      TextureAtlasUv.y + TextureAtlasUv.height - epsilon));
            Uvs.Add(new Vector2(TextureAtlasUv.x + TextureAtlasUv.width - epsilon,
                                      TextureAtlasUv.y + epsilon));
        }

        public virtual Rect GetTextureAtlasUv()
        {
            return Map.Instance.TextureAtlasUvs[2];
        }

        public virtual Texture2D GetMainTexture()
        {
            return Map.Instance.TextureAtlas;
        }
    }

}