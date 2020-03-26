using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using NMObj;

public partial class Map : MonoBehaviour
{

    public Transform ObjPrefab;
    public Transform ObjParent;
    public Transform m_mapCamera;
    private static Map m_instance;
    private Obj m_player = null;
    private Vector3 m_mapCenterPos;
    private Obj[] m_terrains;
    private uint m_serverFrame;
    public static Net.p_obj PObj = null;

    public Obj Player { get { return m_player; } }

    public uint ServerFrame
    {
        set { m_serverFrame = value; }
        get { return m_serverFrame; }
    }

    public Vector3 MapCenterPos
    {
        get { return m_mapCenterPos; }
        set { m_mapCenterPos = value; }
    }

    public static Map Instance
    {
        get { return m_instance; }
        set { m_instance = value; }
    }
    void Awake()
    {
        TextureAtlas = new Texture2D(2048, 2048, TextureFormat.ARGB32, false);
        TextureAtlasUvs = TextureAtlas.PackTextures(Textures, 0);
        TextureAtlas.filterMode = FilterMode.Point;
        TextureAtlas.anisoLevel = 9;
        TextureAtlas.Apply();
        Instance = this;
    }
    // Use this for initialization
    void Start()
    {

    }

    void SetPlayer(Obj playerObj)
    {
        var Listener = gameObject.GetComponent<PlayerListener>();
        Listener.PlayerObj = playerObj;
    }

    // Update is called once per frame
    void Update()
    {
        if (PObj != null && m_player == null)
        {
            m_player = MapMsg.TryAddObj(PObj);

            //m_player.Move(ObjDirection.FORWARD);
            m_mapCamera.transform.localRotation = Quaternion.Euler(45, 0, 0);
            MapCenterPos = m_player.GPose;
            SetPlayer(m_player);
            initTerrain();
        }
        // Debug.Log(m_player == null);
        if (m_player != null)
        {
            updateCamera();
            updateTerrain();
        }


    }

    void updateCamera()
    {
        m_mapCamera.transform.localPosition = m_player.ObjTransform.localPosition + new Vector3(-5.0f, -71.3f, -7f);

    }
    void updateTerrain()
    {
        var Width = Config.TerrainArrayX * Config.TerrainMeshX;
        var Height = Config.TerrainArrayY * Config.TerrainMeshY;
        if (Mathf.Abs(m_player.GPose.x - MapCenterPos.x) > Width / 4)
        {
            int i = 0;
            for (i = 0; i < Config.TerrainArrayX * Config.TerrainArrayY; i++)
            {
                // 离中心点超过1/4 的距离，且和 m_player不在同一边
                if (Mathf.Abs(m_terrains[i].GPose.x - MapCenterPos.x) >= Width / 4 &&
                    Mathf.Abs(m_terrains[i].GPose.x - m_player.GPose.x) >= Width / 2)
                {
                    var index = 1;
                    if (m_terrains[i].GPose.x > MapCenterPos.x)
                    {
                        index = -1;
                    }
                    Vector3 OffPos = new Vector3(index * Width, 0, 0); //todo 按理说应该是移动3/4距离的，但是移动3/4效果不对，移动4/4才对
                    m_terrains[i].GPose += OffPos;
                    m_terrains[i].ObjTransform.position += OffPos;
                    return;
                }
            }
            if (i == Config.TerrainArrayX * Config.TerrainArrayY)
            {
                // 中心点 向左或向右1/4的距离
                if (MapCenterPos.x > m_player.GPose.x)
                {
                    MapCenterPos -= new Vector3(Width / 4, 0, 0);
                }
                else
                {
                    MapCenterPos += new Vector3(Width / 4, 0, 0);
                }
            }

        }
        else if (Mathf.Abs(m_player.GPose.z - MapCenterPos.z) > Height / 4)
        {
            int i = 0;
            for (i = 0; i < Config.TerrainArrayX * Config.TerrainArrayY; i++)
            {
                if (Mathf.Abs(m_terrains[i].GPose.z - MapCenterPos.z) >= Height / 4 &&
                    Mathf.Abs(m_terrains[i].GPose.z - m_player.GPose.z) >= Height / 2)
                {
                    var index = 1;
                    if (m_terrains[i].GPose.z > MapCenterPos.z)
                    {
                        index = -1;
                    }
                    Vector3 OffPos = new Vector3(0, 0, index * Height);
                    m_terrains[i].GPose += OffPos;
                    m_terrains[i].ObjTransform.position += OffPos;
                    return;
                }
            }
            if (i == Config.TerrainArrayX * Config.TerrainArrayY)
            {
                // 中心点 向前或向后1/4的距离
                if (MapCenterPos.z > m_player.GPose.z)
                {
                    MapCenterPos -= new Vector3(0, 0, Height / 4);
                }
                else
                {
                    MapCenterPos += new Vector3(0, 0, Height / 4);
                }
            }
        }
    }

    void initTerrain()
    {
        m_terrains = new Obj[Config.TerrainArrayX * Config.TerrainArrayY];
        for (int x = -Config.TerrainArrayX / 2; x < Config.TerrainArrayX / 2; x++)
        {
            for (int z = -Config.TerrainArrayY / 2; z < Config.TerrainArrayY / 2; z++)
            {
                int i = x + Config.TerrainArrayX / 2, j = z + Config.TerrainArrayY / 2;
                Obj terrain = ObjMgr.NewObj(new Vector3(x * Config.TerrainMeshX + MapCenterPos.x, 0, z * Config.TerrainMeshY + MapCenterPos.z), T_Obj.T_TERRAIN.GenObjType(), i * Config.TerrainArrayY + j+1000000);
                terrain.Init();
                m_terrains[i * Config.TerrainArrayY + j] = terrain;
            }
        }
    }
}
