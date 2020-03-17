using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CameraMove : MonoBehaviour {
	//player info
	public GameObject player;
	private Transform playerTF;

	//distance vector between player and camera
	private Vector3 dirVector;
	public float distance;

	//mouse move
	private float fMouseX;
	private float fMouseY;
	public float speed;
	public float bottomLimitAngle;//the limit angle
	private float bottomLimit;//the cos value

	// Use this for initialization
	void Start () {
		//**System settings**
		Cursor.visible = false;
		Cursor.lockState = CursorLockMode.Locked;

		//**assignment initial**

		//player assignment
		//player = GameObject.FindGameObjectWithTag ("Player");
		playerTF = player.transform;


		//**value initial**

		//distanceVector initial
		dirVector = Vector3.Normalize(playerTF.position - transform.position);
		transform.position = playerTF.position + distance * (-dirVector);

		//mouse move initial
		fMouseX = 0;
		fMouseY = 0;
		bottomLimit = Mathf.Cos (bottomLimitAngle / 180 * Mathf.PI);
	}
    int index = 0;
	// Update is called once per frame
	void Update () {
		//Update Camera
		transform.LookAt (playerTF);
		//transform.position = playerTF.position - distanceVector;

		//Camera Move
		fMouseX = Input.GetAxis ("Mouse X");
		fMouseY = Input.GetAxis ("Mouse Y");
       
        //avoid dithering
        if (Vector3.Dot (-dirVector.normalized, -playerTF.up.normalized) > bottomLimit) {
			if (fMouseY > 0) {
				fMouseY = 0;
			};
		}

		//two types of parameters;
		//(axis,value)is rotate around the axis of the transform's position;
		// (position, axis, value)is rotate around the axis of the specific position;
		//Rotate Horizontal
		transform.RotateAround(playerTF.position ,playerTF.up, speed * fMouseX);
		//Rotate Vertical
		transform.RotateAround (playerTF.position, -VerticalRotateAxis(dirVector),speed * fMouseY);

		//distance Control
		dirVector = Vector3.Normalize(playerTF.position - transform.position);
		Ray cameraRay = new Ray(playerTF.position, -dirVector);
		RaycastHit hitinfo;
		if (Physics.Raycast (cameraRay, out hitinfo, distance, LayerMask.GetMask("Terrain"))) {
			
		} else {
			transform.position = playerTF.position + distance * (-dirVector);
		}

	}

	Vector3 VerticalRotateAxis(Vector3 dirVector){
		Vector3 player2Camera = -dirVector.normalized;
		float x = player2Camera.x;
		float z = player2Camera.z;
		Vector3 rotateAxis = Vector3.zero;
		rotateAxis.z = Mathf.Sqrt (x * x / (x * x + z * z));
		rotateAxis.x = Mathf.Sqrt (z * z / (x * x + z * z));
		if (x >= 0) {
			if (z >= 0) {
				rotateAxis.x = -rotateAxis.x;
			}
		} else {
			if (z >= 0) {
				rotateAxis.x = -rotateAxis.x;
				rotateAxis.z = -rotateAxis.z;
			} else {
				rotateAxis.z = -rotateAxis.z;
			}
		}
		return rotateAxis;
	}
}
